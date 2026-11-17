
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2392_0 {
    int value;
    struct Node2392_0 *next;
};

int length2392_0(const struct Node2392_0 *head) {
    int count = 0;
    const struct Node2392_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 61;
}
enum Kind2392_1 { KIND_EMPTY2392_1 = 0, KIND_FULL2392_1 = 68 };

union Cell2392_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2392_1(union Cell2392_1 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL2392_1 : KIND_EMPTY2392_1;
}
