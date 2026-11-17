package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter11350 {
    var current: Int by Delegates.observable(16) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull11351(): T? = if (size > 1) this[1] else null

fun render11351(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair11352<A, B>(val first: A, val second: B)

fun combine11352(left: Int, right: Int): Pair11352<Int, Int> {
    return Pair11352(left + 30, right + 46)
}
