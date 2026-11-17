
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind985_0 { KIND_EMPTY985_0 = 0, KIND_FULL985_0 = 6 };

union Cell985_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind985_0(union Cell985_0 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL985_0 : KIND_EMPTY985_0;
}
struct Node985_1 {
    int value;
    struct Node985_1 *next;
};

int length985_1(const struct Node985_1 *head) {
    int count = 0;
    const struct Node985_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
