package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull17770(): T? = if (size > 1) this[1] else null

fun render17770(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter17771 {
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
data class Pair17772<A, B>(val first: A, val second: B)

fun combine17772(left: Int, right: Int): Pair17772<Int, Int> {
    return Pair17772(left + 34, right + 26)
}
@JvmName("step17773")
fun step17773(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
sealed class Outcome17774 {
    data class Success(val payload: String) : Outcome17774()
    data object Failure : Outcome17774()
}

fun describe17774(outcome: Outcome17774): String = when (outcome) {
    is Outcome17774.Success -> outcome.payload
    Outcome17774.Failure -> "failed"
}
