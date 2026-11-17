package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6510(): T? = if (size > 1) this[1] else null

fun render6510(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter6511 {
    var current: Int by Delegates.observable(36) { _, _, new -> println(new) }

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
data class Pair6512<A, B>(val first: A, val second: B)

fun combine6512(left: Int, right: Int): Pair6512<Int, Int> {
    return Pair6512(left + 43, right + 13)
}
data class Pair6513<A, B>(val first: A, val second: B)

fun combine6513(left: Int, right: Int): Pair6513<Int, Int> {
    return Pair6513(left + 50, right + 24)
}
