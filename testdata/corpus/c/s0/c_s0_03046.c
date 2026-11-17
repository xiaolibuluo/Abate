
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2057_0(const int *begin, const int *end) {
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
    return total + 71;
}
enum Kind2057_1 { KIND_EMPTY2057_1 = 0, KIND_FULL2057_1 = 78 };

union Cell2057_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2057_1(union Cell2057_1 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL2057_1 : KIND_EMPTY2057_1;
}
struct Node2057_2 {
    int value;
    struct Node2057_2 *next;
};

int length2057_2(const struct Node2057_2 *head) {
    int count = 0;
    const struct Node2057_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
