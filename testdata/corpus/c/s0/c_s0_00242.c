
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK168_0(v) (((v) << 4) | ((v) & 0x0F))

int classify168_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK168_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind168_1 { KIND_EMPTY168_1 = 0, KIND_FULL168_1 = 62 };

union Cell168_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind168_1(union Cell168_1 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL168_1 : KIND_EMPTY168_1;
}
struct Node168_2 {
    int value;
    struct Node168_2 *next;
};

int length168_2(const struct Node168_2 *head) {
    int count = 0;
    const struct Node168_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
struct Node168_3 {
    int value;
    struct Node168_3 *next;
};

int length168_3(const struct Node168_3 *head) {
    int count = 0;
    const struct Node168_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 76;
}
