
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind202_0 { KIND_EMPTY202_0 = 0, KIND_FULL202_0 = 12 };

union Cell202_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind202_0(union Cell202_0 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL202_0 : KIND_EMPTY202_0;
}
struct Node202_1 {
    int value;
    struct Node202_1 *next;
};

int length202_1(const struct Node202_1 *head) {
    int count = 0;
    const struct Node202_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
