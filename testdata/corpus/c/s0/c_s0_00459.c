
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK313_0(v) (((v) << 4) | ((v) & 0x0F))

int classify313_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK313_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan313_1(const int *begin, const int *end) {
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
    return total + 7;
}
enum Kind313_2 { KIND_EMPTY313_2 = 0, KIND_FULL313_2 = 14 };

union Cell313_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind313_2(union Cell313_2 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL313_2 : KIND_EMPTY313_2;
}
struct Node313_3 {
    int value;
    struct Node313_3 *next;
};

int length313_3(const struct Node313_3 *head) {
    int count = 0;
    const struct Node313_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
