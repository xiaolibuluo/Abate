package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4100<A, B>(val first: A, val second: B)

fun combine4100(left: Int, right: Int): Pair4100<Int, Int> {
    return Pair4100(left + 97, right + 12)
}
sealed class Outcome4101 {
    data class Success(val payload: String) : Outcome4101()
    data object Failure : Outcome4101()
}

fun describe4101(outcome: Outcome4101): String = when (outcome) {
    is Outcome4101.Success -> outcome.payload
    Outcome4101.Failure -> "failed"
}
fun <T> List<T>.secondOrNull4102(): T? = if (size > 1) this[1] else null

fun render4102(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter4103 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
