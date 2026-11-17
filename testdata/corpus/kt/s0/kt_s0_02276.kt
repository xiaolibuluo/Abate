package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair18670<A, B>(val first: A, val second: B)

fun combine18670(left: Int, right: Int): Pair18670<Int, Int> {
    return Pair18670(left + 26, right + 24)
}
data class Pair18671<A, B>(val first: A, val second: B)

fun combine18671(left: Int, right: Int): Pair18671<Int, Int> {
    return Pair18671(left + 33, right + 35)
}
class Counter18672 {
    var current: Int by Delegates.observable(40) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull18673(): T? = if (size > 1) this[1] else null

fun render18673(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
