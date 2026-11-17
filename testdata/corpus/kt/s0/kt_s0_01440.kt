package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step11870")
fun step11870(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 13
}
@JvmName("step11871")
fun step11871(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
class Counter11872 {
    var current: Int by Delegates.observable(27) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
@JvmName("step11873")
fun step11873(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 34
}
fun <T> List<T>.secondOrNull11874(): T? = if (size > 1) this[1] else null

fun render11874(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
