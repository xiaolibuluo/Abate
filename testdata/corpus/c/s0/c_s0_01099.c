
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind751_0 { KIND_EMPTY751_0 = 0, KIND_FULL751_0 = 68 };

union Cell751_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind751_0(union Cell751_0 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL751_0 : KIND_EMPTY751_0;
}
struct Node751_1 {
    int value;
    struct Node751_1 *next;
};

int length751_1(const struct Node751_1 *head) {
    int count = 0;
    const struct Node751_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
int scan751_2(const int *begin, const int *end) {
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
    return total + 82;
}
struct Node751_3 {
    int value;
    struct Node751_3 *next;
};

int length751_3(const struct Node751_3 *head) {
    int count = 0;
    const struct Node751_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
