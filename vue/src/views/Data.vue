<template>
  <div style="max-width:100%; text-align:center; margin:auto">
    <MDBAccordion v-model="activeItem" stayOpen>
      <MDBAccordionItem headerTitle="Osmosis" collapseId="stayOpen-collapseOne">
        <tr>
          <td>
            <MDBCard>
              <a v-mdb-ripple="{ color: 'light' }">
                <MDBCardImg src="https://miro.medium.com/max/1400/1*_iE5yi1KqS8nwijja0vEQg.jpeg" top alt="..."
                  style="width:300px;" />
              </a>
              <MDBCardBody>
                <MDBCardTitle>Bridge</MDBCardTitle>
                <MDBCardText>
                  This controls the bridiging functionality between Osmosis and its IBC connected chains. You can pause,
                  shutdown, or send alerts to the bridge..
                </MDBCardText>
                <MDBBtn tag="a" href="#!" color="primary">Pause Bridge</MDBBtn>
                <MDBBtn tag="a" href="#!" color="primary">Shutdown Bridge</MDBBtn>
                <MDBBtn tag="a" href="#!" color="primary">Send Bridge Alert</MDBBtn>
              </MDBCardBody>
            </MDBCard>
          </td>
          <td>
            <MDBCard>
              <a v-mdb-ripple="{ color: 'light' }">
                <MDBCardImg src="https://i.ytimg.com/vi/0gcGrvZ5wFw/maxresdefault.jpg" top alt="..."
                  style="width:500px;" />
              </a>
              <MDBCardBody>
                <MDBCardTitle>Swaps</MDBCardTitle>
                <MDBCardText>
                  This controls the swap functionality between Osmosis and its IBC connected chains. You can pause,
                  shutdown, or send alerts to the swap..
                </MDBCardText>
                <MDBBtn tag="a" href="#!" color="primary">Pause Swaps</MDBBtn>
                <MDBBtn tag="a" href="#!" color="primary">Shutdown Swaps</MDBBtn>
                <MDBBtn tag="a" href="#!" color="primary">Send Swap Alert</MDBBtn>
              </MDBCardBody>
            </MDBCard>
          </td>
          <td>
            <MDBCard>
              <a v-mdb-ripple="{ color: 'light' }">
                <MDBCardImg src="https://i.ytimg.com/vi/q_hyGhGkFSc/maxresdefault.jpg" top alt="..."
                  style="width:500px;" />
              </a>
              <MDBCardBody>
                <MDBCardTitle>Bridge</MDBCardTitle>
                <MDBCardText>
                  This controls the bridiging functionality between Osmosis and its IBC connected chains. You can pause,
                  shutdown, or send alerts to the bridge..
                </MDBCardText>
                <MDBBtn tag="a" href="#!" color="primary">Pause Bridge</MDBBtn>
                <MDBBtn tag="a" href="#!" color="primary">Shutdown Bridge</MDBBtn>
                <MDBBtn tag="a" href="#!" color="primary">Send Bridge Alert</MDBBtn>
              </MDBCardBody>
            </MDBCard>
          </td>
        </tr>
      </MDBAccordionItem>
      <MDBAccordionItem headerTitle="Evmos" collapseId="stayOpen-collapseOne">
        <MDBCard>
          <a v-mdb-ripple="{ color: 'light' }">
            <MDBCardImg src="https://mdbootstrap.com/img/new/standard/nature/184.webp" top alt="..." />
          </a>
          <MDBCardBody>
            <MDBCardTitle>Card title</MDBCardTitle>
            <MDBCardText>
              Some quick example text to build on the card title and make up the bulk of the
              card's content.
            </MDBCardText>
            <MDBBtn tag="a" href="#!" color="primary">Button</MDBBtn>
          </MDBCardBody>
        </MDBCard>
      </MDBAccordionItem>
      <MDBAccordionItem headerTitle="Terra" collapseId="stayOpen-collapseOne">
        <MDBCard>
          <a v-mdb-ripple="{ color: 'light' }">
            <MDBCardImg src="https://mdbootstrap.com/img/new/standard/nature/184.webp" top alt="..." />
          </a>
          <MDBCardBody>
            <MDBCardTitle>Card title</MDBCardTitle>
            <MDBCardText>
              Some quick example text to build on the card title and make up the bulk of the
              card's content.
            </MDBCardText>
            <MDBBtn tag="a" href="#!" color="primary">Button</MDBBtn>
          </MDBCardBody>
        </MDBCard>
      </MDBAccordionItem>
    </MDBAccordion>
  </div>
</template>

<script>
import { MDBAccordion, MDBAccordionItem } from "mdb-vue-ui-kit";
import { MDBCard, MDBCardBody, MDBCardTitle, MDBCardText, MDBCardImg, MDBBtn, mdbRipple } from "mdb-vue-ui-kit";
import { ref } from "vue";

export default {
  name: "Data",
  components: {
    MDBAccordion,
    MDBAccordionItem,
    MDBCard,
    MDBCardBody,
    MDBCardTitle,
    MDBCardText,
    MDBCardImg,
    MDBBtn,
  },
  setup() {
    const activeItem = ref('stayOpen-collapseOne');
    return {
      activeItem
    }
  },
  methods: {
    makeToast(variant = null) {
      this.$bvToast.toast('Toast body content', {
        title: `Variant ${variant || 'default'}`,
        variant: variant,
        solid: true
      })
    }
  }
};
</script>
