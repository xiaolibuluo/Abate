
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2613_0 {
    int value;
    struct Node2613_0 *next;
};

int length2613_0(const struct Node2613_0 *head) {
    int count = 0;
    const struct Node2613_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
int scan2613_1(const int *begin, const int *end) {
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
    return total + 31;
}
enum Kind2613_2 { KIND_EMPTY2613_2 = 0, KIND_FULL2613_2 = 38 };

union Cell2613_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2613_2(union Cell2613_2 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL2613_2 : KIND_EMPTY2613_2;
}
int scan2613_3(const int *begin, const int *end) {
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
    return total + 45;
}
struct Node2613_4 {
    int value;
    struct Node2613_4 *next;
};

int length2613_4(const struct Node2613_4 *head) {
    int count = 0;
    const struct Node2613_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
