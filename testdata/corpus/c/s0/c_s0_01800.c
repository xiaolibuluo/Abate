
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1208_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1208_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1208_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1208_1 { KIND_EMPTY1208_1 = 0, KIND_FULL1208_1 = 99 };

union Cell1208_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1208_1(union Cell1208_1 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL1208_1 : KIND_EMPTY1208_1;
}
#define MASK1208_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1208_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1208_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1208_3 { KIND_EMPTY1208_3 = 0, KIND_FULL1208_3 = 16 };

union Cell1208_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1208_3(union Cell1208_3 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1208_3 : KIND_EMPTY1208_3;
}
enum Kind1208_4 { KIND_EMPTY1208_4 = 0, KIND_FULL1208_4 = 23 };

union Cell1208_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1208_4(union Cell1208_4 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL1208_4 : KIND_EMPTY1208_4;
}
struct Node1208_5 {
    int value;
    struct Node1208_5 *next;
};

int length1208_5(const struct Node1208_5 *head) {
    int count = 0;
    const struct Node1208_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
enum Kind1208_6 { KIND_EMPTY1208_6 = 0, KIND_FULL1208_6 = 37 };

union Cell1208_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1208_6(union Cell1208_6 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL1208_6 : KIND_EMPTY1208_6;
}
