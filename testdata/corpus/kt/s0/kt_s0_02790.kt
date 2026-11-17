package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23040<A, B>(val first: A, val second: B)

fun combine23040(left: Int, right: Int): Pair23040<Int, Int> {
    return Pair23040(left + 81, right + 33)
}
fun <T> List<T>.secondOrNull23041(): T? = if (size > 1) this[1] else null

fun render23041(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome23042 {
    data class Success(val payload: String) : Outcome23042()
    data object Failure : Outcome23042()
}

fun describe23042(outcome: Outcome23042): String = when (outcome) {
    is Outcome23042.Success -> outcome.payload
    Outcome23042.Failure -> "failed"
}
class Counter23043 {
    var current: Int by Delegates.observable(5) { _, _, new -> println(new) }

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
