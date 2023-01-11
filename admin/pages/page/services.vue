<template>
  <div class="about">
    <custom-container @save="save" @edit="toggleEdit" @cancel="cancel" :hideCreate="true">
      <custom-box :title="'Section: Page Header'">
        <image-upload
          label="Header Image"
          :readonly="EditMode"
          v-model="pageData.attributes.header.image"
          @image="isImageValid($event)"
        />
      </custom-box>
      <custom-box :title="'Section: 2'">
        <template v-slot:button></template>
        <b-field
          label="Title"
          custom-class="is-small"
          :type="errors.has('title') ? 'is-danger' : ''"
          :message="errors.first('title')"
        >
          <b-input
            expanded
            v-model="pageData.attributes.content.title"
            name="title"
            :disabled="EditMode"
            placeholder="Title"
            v-validate="'required'"
          />
        </b-field>
        <b-field
          label="Content"
          custom-class="is-small"
          :type="errors.has('content') ? 'is-danger' : ''"
          :message="errors.first('content')"
        >
          <custom-editor
            :class="errors.has('title') ? 'custom-help' : ''"
            name="content"
            :disabled="EditMode"
            v-validate="'required'"
            v-model="pageData.attributes.content.content"
          ></custom-editor>
        </b-field>
        <!-- <b-table
          :data="pageData.attributes.content.List"
          v-if="pageData.attributes.content.List.length != 0"
        >
          <template slot-scope="props">
            <b-table-column label="Content">
              {{ props.row.content }}
            </b-table-column>

            <b-table-column label=" " class="has-text-right" v-if="!EditMode">
              <a
                @click="editService(props.row)"
                class="mdi mdi-pencil icobtn"
              ></a>

              <a
                @click="deleteService(props.index)"
                class="mdi mdi-delete icobtn"
              />
            </b-table-column>
          </template>

          <template slot="empty">
            <section class="section">
              <div class="content has-text-grey has-text-centered">
                <p>
                  <i class="mdi mdi-information-outline mdi-48px" />
                </p>
                <p class="is-size-5">No record available</p>
              </div>
            </section>
          </template>
        </b-table>-->
      </custom-box>
    </custom-container>
  </div>
</template>

<script>
import PagesMix from '@/mixins/admin_pages'
import CustomEditor from '@/components/widgets/CustomEditor'
import ServiceEditor from '@/components/widgets/services/EditorBlock'
import { copy, clone, Notification } from '@/utils/helpers'

export default {
  mixins: [PagesMix],
  components: {
    CustomEditor,
    ServiceEditor
  },
  methods: {
    deleteService() {},
    cancel() {
      this.$router.push('/page')
    },
    addBlock() {
      let block = {
        title: 'Affordable',
        content:
          'Prudent Domiciliary Care are designed to meet all your needs in a very special and personalized approach. We believe that Quality Home Care should be available to everyone. You or your loved ones deserve the home care which is meeting all your needs and being cost effective at the same time'
      }
    },
    editService() {
      let that = this
      this.$buefy.modal.open({
        parent: this,
        component: ServiceEditor,
        hasModalCard: true,
        trapFocus: true,
        events: {
          input(val) {}
        }
      })
    },
    async save() {
      let result = await this.$validator.validateAll()
      if (result) {
        let form = copy(this.pageData)
        //         form.attributes.content.content = `<ul style="list-style-type: disc;">
        // <li>&nbsp;Assist with a &lsquo;good morning&rsquo; start to the day, helping you to get up, wash, shower or bathe, dress and have breakfast</li>
        // <li>Remind or assist you to take your medicines and collect or return medication from your pharmacy or dispensing GP surgery</li>
        // <li>Prepare meals with or for you and assist you at mealtimes</li>
        // <li>Collect your pension for you or with you</li>
        // <li>Shop with you or help you to make a shopping list, go to the shops, come back and put it all away</li>
        // <li>Help with your household tasks such as laundry or ironing and keeping your home clean and tidy</li>
        // <li>Support you with social activities such as going out for a walk, attending a day centre, visiting friends or family, going to your church or club etc</li>
        // <li>Give a little pampering when you need it &ndash; the occasional visit to organise bathing, clean clothes, fresh bedding and a thoughtfully prepared meal before bedtime</li>
        // <li>At the end of each day, some help with getting ready for bed</li>
        // </ul>
        // <p><strong>Our Services also includes</strong></p>
        // <ul style="list-style-type: disc;">
        // <li><a href="https://pbgcare.co.uk/services/Companionship-care">Companionship Care</a></li>
        // <li><a href="https://pbgcare.co.uk/services/Palliative-care">Palliative Care</a></li>
        // <li><a href="https://pbgcare.co.uk/services/emphatic-end-of-life-care">Emphatic End Of Life Care</a></li>
        // <li><a href="https://pbgcare.co.uk/services/long-term-home-care">Long Term Home Care</a></li>
        // <li><a href="https://pbgcare.co.uk/services/dementia-care-at-home">Dementia Care At Home</a></li>
        // <li><a href="https://pbgcare.co.uk/services/home-from-hospital">Home From Hospital</a></li>
        // <li><a href="https://pbgcare.co.uk/services/take-a-break">Take A Break</a></li>
        // <li><a href="https://pbgcare.co.uk/services/live-in-care">Live In Care</a></li>
        // <li><a href="https://pbgcare.co.uk/services/Physical-disability-Support-at-home">Physical Disability Support At Home</a></li>
        // <li><a href="https://pbgcare.co.uk/services/supporting-independent-living">Supporting Independent Living</a></li>
        // <li><a href="https://pbgcare.co.uk/services/learning-disabilities">Learning Disabilities</a></li>
        // </ul>`

        let param = {
          url: `pages/${this.uuid}`,
          data: form
        }
        let req = await this.$store.dispatch('Postdata', param)
        if (req.done) {
          Notification(this, 'Changes saved successfully', 'is-success')
          this.$router.push('/page')
        } else {
          Notification(this, 'an error occured', 'is-danger')
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped></style>
