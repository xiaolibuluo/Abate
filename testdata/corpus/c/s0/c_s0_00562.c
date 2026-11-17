
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node385_0 {
    int value;
    struct Node385_0 *next;
};

int length385_0(const struct Node385_0 *head) {
    int count = 0;
    const struct Node385_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
enum Kind385_1 { KIND_EMPTY385_1 = 0, KIND_FULL385_1 = 70 };

union Cell385_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind385_1(union Cell385_1 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL385_1 : KIND_EMPTY385_1;
}
int scan385_2(const int *begin, const int *end) {
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
    return total + 77;
}
#define MASK385_3(v) (((v) << 4) | ((v) & 0x0F))

int classify385_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK385_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan385_4(const int *begin, const int *end) {
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
    return total + 91;
}
#define MASK385_5(v) (((v) << 4) | ((v) & 0x0F))

int classify385_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK385_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan385_6(const int *begin, const int *end) {
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
    return total + 8;
}
