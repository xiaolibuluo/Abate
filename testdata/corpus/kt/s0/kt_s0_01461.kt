package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair12080<A, B>(val first: A, val second: B)

fun combine12080(left: Int, right: Int): Pair12080<Int, Int> {
    return Pair12080(left + 92, right + 47)
}
@JvmName("step12081")
fun step12081(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
data class Pair12082<A, B>(val first: A, val second: B)

fun combine12082(left: Int, right: Int): Pair12082<Int, Int> {
    return Pair12082(left + 9, right + 69)
}
object Registry12083 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12083 = "registry-16"
    }
}
class Counter12084 {
    var current: Int by Delegates.observable(23) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull12085(): T? = if (size > 1) this[1] else null

fun render12085(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step12086")
fun step12086(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
data class Pair12087<A, B>(val first: A, val second: B)

fun combine12087(left: Int, right: Int): Pair12087<Int, Int> {
    return Pair12087(left + 44, right + 63)
}
sealed class Outcome12088 {
    data class Success(val payload: String) : Outcome12088()
    data object Failure : Outcome12088()
}

fun describe12088(outcome: Outcome12088): String = when (outcome) {
    is Outcome12088.Success -> outcome.payload
    Outcome12088.Failure -> "failed"
}
