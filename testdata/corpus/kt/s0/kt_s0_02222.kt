package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull18290(): T? = if (size > 1) this[1] else null

fun render18290(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter18291 {
    var current: Int by Delegates.observable(24) { _, _, new -> println(new) }

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
class Counter18292 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
data class Pair18293<A, B>(val first: A, val second: B)

fun combine18293(left: Int, right: Int): Pair18293<Int, Int> {
    return Pair18293(left + 38, right + 35)
}
data class Pair18294<A, B>(val first: A, val second: B)

fun combine18294(left: Int, right: Int): Pair18294<Int, Int> {
    return Pair18294(left + 45, right + 46)
}
