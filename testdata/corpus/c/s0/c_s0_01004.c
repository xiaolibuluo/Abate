
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node691_0 {
    int value;
    struct Node691_0 *next;
};

int length691_0(const struct Node691_0 *head) {
    int count = 0;
    const struct Node691_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 64;
}
enum Kind691_1 { KIND_EMPTY691_1 = 0, KIND_FULL691_1 = 71 };

union Cell691_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind691_1(union Cell691_1 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL691_1 : KIND_EMPTY691_1;
}
