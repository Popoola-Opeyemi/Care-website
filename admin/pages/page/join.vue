<template>
  <div>
    <custom-container @save="save" @edit="toggleEdit" @cancel="cancel" :hideCreate="true">
      <custom-box :title="'Section: Page Header'">
        <image-upload
          label="Header Image"
          :readonly="EditMode"
          v-model="pageData.attributes.header.image"
          @image="isImageValid($event)"
        />
        <b-field
          label="Title"
          custom-class="is-small"
          :type="errors.has('title') ? 'is-danger': ''"
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
      </custom-box>
    </custom-container>
  </div>
</template>

<script>
import PagesMix from '@/mixins/admin_pages'
import CustomEditor from '@/components/widgets/CustomEditor'
import CustomDropdown from '@/components/widgets/process/Dropdown'

import ServiceEditor from '@/components/widgets/services/EditorBlock'
import { copy, clone, Notification } from '@/utils/helpers'

export default {
  mixins: [PagesMix],
  components: {
    CustomEditor,
    ServiceEditor,
    CustomDropdown
  },
  methods: {
    cancel() {
      this.$router.push('/page')
    },
    async save() {
      let result = await this.$validator.validateAll()
      if (result) {
        let form = copy(this.pageData)
        //         form.attributes.content.content = `<p><strong>If you are seeking a role with real job satisfaction, join us now and benefit from our excellent rates of pay of &pound;9.03 (flat rate) - &pound;9.45 (unsociable hour rate) per hour.</strong></p>
        // <p>&nbsp;</p>
        // <p>Prudent Domiciliary Care is the leading provider of home care in Kent and Royal Borough of Greenwich, providing services throughout Charlton, Plumstead and Abbey Wood, Belvedere, Sidcup. we provide high quality, professional care to people in their own homes including the following:</p>
        // <p>&nbsp;</p>
        // <p><strong>Responsibilities</strong></p>
        // <ul>
        // <li>Personal care and helping clients to wash and dress</li>
        // <li>Assistance with medication</li>
        // <li>Shopping and support with everyday tasks</li>
        // <li>Light household chores</li>
        // <li>Meal preparation</li>
        // </ul>
        // <p>&nbsp;</p>
        // <p><strong>Requirements</strong></p>
        // <p>Whilst previous care experience is helpful, it is not essential, as we will provide you with the training you require to carry out your role with confidence and a support team who you can access at any time. We are seeking people who are:</p>
        // <ul>
        // <li>Kind, caring and wish to make a difference</li>
        // <li>Reliable and take pride in their work</li>
        // <li>Have a full UK driving licence and own a vehicle for use at work</li>
        // <li>Flexible in their approach and able to work alternate weekends</li>
        // <li>Available to take part in our introductory 4 day training programme</li>
        // </ul>
        // <p>&nbsp;</p>
        // <p>All Applications should be sent to&nbsp;<a  href="mailto:info@pbgcare.co.uk" "">info@pbgcare.co.uk</a></p>`
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
