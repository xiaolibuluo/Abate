package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9950<A, B>(val first: A, val second: B)

fun combine9950(left: Int, right: Int): Pair9950<Int, Int> {
    return Pair9950(left + 39, right + 20)
}
fun <T> List<T>.secondOrNull9951(): T? = if (size > 1) this[1] else null

fun render9951(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome9952 {
    data class Success(val payload: String) : Outcome9952()
    data object Failure : Outcome9952()
}

fun describe9952(outcome: Outcome9952): String = when (outcome) {
    is Outcome9952.Success -> outcome.payload
    Outcome9952.Failure -> "failed"
}
class Counter9953 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
