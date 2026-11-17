
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind756_0 { KIND_EMPTY756_0 = 0, KIND_FULL756_0 = 36 };

union Cell756_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind756_0(union Cell756_0 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL756_0 : KIND_EMPTY756_0;
}
#define MASK756_1(v) (((v) << 4) | ((v) & 0x0F))

int classify756_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK756_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node756_2 {
    int value;
    struct Node756_2 *next;
};

int length756_2(const struct Node756_2 *head) {
    int count = 0;
    const struct Node756_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
int scan756_3(const int *begin, const int *end) {
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
    return total + 57;
}
