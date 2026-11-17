package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14060 {
    data class Success(val payload: String) : Outcome14060()
    data object Failure : Outcome14060()
}

fun describe14060(outcome: Outcome14060): String = when (outcome) {
    is Outcome14060.Success -> outcome.payload
    Outcome14060.Failure -> "failed"
}
fun <T> List<T>.secondOrNull14061(): T? = if (size > 1) this[1] else null

fun render14061(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step14062")
fun step14062(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
sealed class Outcome14063 {
    data class Success(val payload: String) : Outcome14063()
    data object Failure : Outcome14063()
}

fun describe14063(outcome: Outcome14063): String = when (outcome) {
    is Outcome14063.Success -> outcome.payload
    Outcome14063.Failure -> "failed"
}
@JvmName("step14064")
fun step14064(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
data class Pair14065<A, B>(val first: A, val second: B)

fun combine14065(left: Int, right: Int): Pair14065<Int, Int> {
    return Pair14065(left + 82, right + 24)
}
class Counter14066 {
    var current: Int by Delegates.observable(89) { _, _, new -> println(new) }

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
@JvmName("step14067")
fun step14067(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
