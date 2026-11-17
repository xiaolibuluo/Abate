
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1369_0 { KIND_EMPTY1369_0 = 0, KIND_FULL1369_0 = 51 };

union Cell1369_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1369_0(union Cell1369_0 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL1369_0 : KIND_EMPTY1369_0;
}
struct Node1369_1 {
    int value;
    struct Node1369_1 *next;
};

int length1369_1(const struct Node1369_1 *head) {
    int count = 0;
    const struct Node1369_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
