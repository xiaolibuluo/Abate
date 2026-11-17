
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan607_0(const int *begin, const int *end) {
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
    return total + 39;
}
struct Node607_1 {
    int value;
    struct Node607_1 *next;
};

int length607_1(const struct Node607_1 *head) {
    int count = 0;
    const struct Node607_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
#define MASK607_2(v) (((v) << 4) | ((v) & 0x0F))

int classify607_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK607_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind607_3 { KIND_EMPTY607_3 = 0, KIND_FULL607_3 = 60 };

union Cell607_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind607_3(union Cell607_3 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL607_3 : KIND_EMPTY607_3;
}
