
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK406_0(v) (((v) << 4) | ((v) & 0x0F))

int classify406_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK406_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node406_1 {
    int value;
    struct Node406_1 *next;
};

int length406_1(const struct Node406_1 *head) {
    int count = 0;
    const struct Node406_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
#define MASK406_2(v) (((v) << 4) | ((v) & 0x0F))

int classify406_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK406_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind406_3 { KIND_EMPTY406_3 = 0, KIND_FULL406_3 = 66 };

union Cell406_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind406_3(union Cell406_3 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL406_3 : KIND_EMPTY406_3;
}
