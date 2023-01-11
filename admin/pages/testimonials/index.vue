<template>
  <div>
    <custom-container :croute="croute" :hide="false">
      <b-table :data="pageList.items">
        <template slot-scope="props">
          <b-table-column label="Date" width="50">
            <nuxt-link
              :to="'/testimonials/' + props.row.id"
            >{{ props.row.created_at | moment('DD/MM/YYYY') }}</nuxt-link>
          </b-table-column>

          <b-table-column label="Author" width="150">
            <nuxt-link :to="'/testimonials/' + props.row.id">
              {{
              props.row.author
              }}
            </nuxt-link>
          </b-table-column>

          <b-table-column label=" " class="has-text-right">
            <nuxt-link :to="'/testimonials/' + props.row.id" class="mdi mdi-eye icobtn" />

            <a @click="deleteAttrib(props.row.id)" class="mdi mdi-delete icobtn" />
          </b-table-column>
        </template>

        <template slot="empty">
          <section class="section">
            <div class="content has-text-grey has-text-centered">
              <p>
                <i class="mdi mdi-information-outline mdi-48px" />
              </p>
              <p class="is-size-5">No Record Available</p>
            </div>
          </section>
        </template>
      </b-table>
      <section class="section">
        <b-pagination
          :total="total"
          @change="setPage($event)"
          :per-page="perPage"
          :current.sync="current"
        />
      </section>
    </custom-container>
  </div>
</template>

<script>
import CustomDropdown from '@/components/widgets/Dropdown'
import PagesMix from '@/mixins/page.index'

export default {
  mixins: [PagesMix],
  components: {
    CustomDropdown
  },
  methods: {}
}
</script>

<style lang="scss" scoped></style>
