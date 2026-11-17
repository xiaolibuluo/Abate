package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9570<A, B>(val first: A, val second: B)

fun combine9570(left: Int, right: Int): Pair9570<Int, Int> {
    return Pair9570(left + 30, right + 59)
}
data class Pair9571<A, B>(val first: A, val second: B)

fun combine9571(left: Int, right: Int): Pair9571<Int, Int> {
    return Pair9571(left + 37, right + 9)
}
sealed class Outcome9572 {
    data class Success(val payload: String) : Outcome9572()
    data object Failure : Outcome9572()
}

fun describe9572(outcome: Outcome9572): String = when (outcome) {
    is Outcome9572.Success -> outcome.payload
    Outcome9572.Failure -> "failed"
}
fun <T> List<T>.secondOrNull9573(): T? = if (size > 1) this[1] else null

fun render9573(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair9574<A, B>(val first: A, val second: B)

fun combine9574(left: Int, right: Int): Pair9574<Int, Int> {
    return Pair9574(left + 58, right + 42)
}
fun <T, R> List<T>.mapChecked9575(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter9576 {
    var current: Int by Delegates.observable(72) { _, _, new -> println(new) }

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
