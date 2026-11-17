
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1230_0 { KIND_EMPTY1230_0 = 0, KIND_FULL1230_0 = 87 };

union Cell1230_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1230_0(union Cell1230_0 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL1230_0 : KIND_EMPTY1230_0;
}
struct Node1230_1 {
    int value;
    struct Node1230_1 *next;
};

int length1230_1(const struct Node1230_1 *head) {
    int count = 0;
    const struct Node1230_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
struct Node1230_2 {
    int value;
    struct Node1230_2 *next;
};

int length1230_2(const struct Node1230_2 *head) {
    int count = 0;
    const struct Node1230_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
