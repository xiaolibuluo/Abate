
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2056_0 {
    int value;
    struct Node2056_0 *next;
};

int length2056_0(const struct Node2056_0 *head) {
    int count = 0;
    const struct Node2056_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
enum Kind2056_1 { KIND_EMPTY2056_1 = 0, KIND_FULL2056_1 = 65 };

union Cell2056_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2056_1(union Cell2056_1 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL2056_1 : KIND_EMPTY2056_1;
}
int scan2056_2(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 72;
}
