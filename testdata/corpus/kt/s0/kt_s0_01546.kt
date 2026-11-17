package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter12800 {
    var current: Int by Delegates.observable(58) { _, _, new -> println(new) }

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
data class Pair12801<A, B>(val first: A, val second: B)

fun combine12801(left: Int, right: Int): Pair12801<Int, Int> {
    return Pair12801(left + 65, right + 13)
}
fun <T> List<T>.secondOrNull12802(): T? = if (size > 1) this[1] else null

fun render12802(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
