import Vue from "vue";
import VueRouter from "vue-router";
import About from "../views/About";
import Home from "../views/Home";
import ServiceTable from "../views/ServiceTable";
import Lambda from "../views/SinglePages/Lambda";
import DynamoDB from "../views/SinglePages/DynamoDB";
import APIGateway from "../views/SinglePages/APIGateway";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
  },
  {
    path: "/about",
    name: "about",
    component: About,
  },
  {
    path: "/service/:serviceName",
    name: "service-table",
    component: ServiceTable,
    props: true,
  },
  {
    path: "/service/lambda/:id",
    name: "lambda",
    component: Lambda,
    props: true,
  },
  {
    path: "/service/dynamodb/:id",
    name: "dynamodb",
    component: DynamoDB,
    props: true,
  },
  {
    path: "/service/apigw/:id",
    name: "apigw",
    component: APIGateway,
    props: true,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
