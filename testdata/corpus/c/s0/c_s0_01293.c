
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind884_0 { KIND_EMPTY884_0 = 0, KIND_FULL884_0 = 51 };

union Cell884_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind884_0(union Cell884_0 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL884_0 : KIND_EMPTY884_0;
}
struct Node884_1 {
    int value;
    struct Node884_1 *next;
};

int length884_1(const struct Node884_1 *head) {
    int count = 0;
    const struct Node884_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
