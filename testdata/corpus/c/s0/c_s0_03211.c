
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2175_0 { KIND_EMPTY2175_0 = 0, KIND_FULL2175_0 = 53 };

union Cell2175_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2175_0(union Cell2175_0 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL2175_0 : KIND_EMPTY2175_0;
}
struct Node2175_1 {
    int value;
    struct Node2175_1 *next;
};

int length2175_1(const struct Node2175_1 *head) {
    int count = 0;
    const struct Node2175_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
enum Kind2175_2 { KIND_EMPTY2175_2 = 0, KIND_FULL2175_2 = 67 };

union Cell2175_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2175_2(union Cell2175_2 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL2175_2 : KIND_EMPTY2175_2;
}
struct Node2175_3 {
    int value;
    struct Node2175_3 *next;
};

int length2175_3(const struct Node2175_3 *head) {
    int count = 0;
    const struct Node2175_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 74;
}
struct Node2175_4 {
    int value;
    struct Node2175_4 *next;
};

int length2175_4(const struct Node2175_4 *head) {
    int count = 0;
    const struct Node2175_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 81;
}
