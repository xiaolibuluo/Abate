
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind930_0 { KIND_EMPTY930_0 = 0, KIND_FULL930_0 = 67 };

union Cell930_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind930_0(union Cell930_0 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL930_0 : KIND_EMPTY930_0;
}
struct Node930_1 {
    int value;
    struct Node930_1 *next;
};

int length930_1(const struct Node930_1 *head) {
    int count = 0;
    const struct Node930_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 74;
}
