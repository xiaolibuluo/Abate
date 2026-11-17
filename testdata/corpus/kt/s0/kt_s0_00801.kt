package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked6590(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull6591(): T? = if (size > 1) this[1] else null

fun render6591(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair6592<A, B>(val first: A, val second: B)

fun combine6592(left: Int, right: Int): Pair6592<Int, Int> {
    return Pair6592(left + 50, right + 69)
}
class Counter6593 {
    var current: Int by Delegates.observable(57) { _, _, new -> println(new) }

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
sealed class Outcome6594 {
    data class Success(val payload: String) : Outcome6594()
    data object Failure : Outcome6594()
}

fun describe6594(outcome: Outcome6594): String = when (outcome) {
    is Outcome6594.Success -> outcome.payload
    Outcome6594.Failure -> "failed"
}
