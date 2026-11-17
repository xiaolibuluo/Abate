
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node314_0 {
    int value;
    struct Node314_0 *next;
};

int length314_0(const struct Node314_0 *head) {
    int count = 0;
    const struct Node314_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
enum Kind314_1 { KIND_EMPTY314_1 = 0, KIND_FULL314_1 = 20 };

union Cell314_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind314_1(union Cell314_1 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL314_1 : KIND_EMPTY314_1;
}
#define MASK314_2(v) (((v) << 4) | ((v) & 0x0F))

int classify314_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK314_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan314_3(const int *begin, const int *end) {
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
    return total + 34;
}
