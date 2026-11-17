package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6600(): T? = if (size > 1) this[1] else null

fun render6600(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull6601(): T? = if (size > 1) this[1] else null

fun render6601(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter6602 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
data class Pair6603<A, B>(val first: A, val second: B)

fun combine6603(left: Int, right: Int): Pair6603<Int, Int> {
    return Pair6603(left + 70, right + 26)
}
sealed class Outcome6604 {
    data class Success(val payload: String) : Outcome6604()
    data object Failure : Outcome6604()
}

fun describe6604(outcome: Outcome6604): String = when (outcome) {
    is Outcome6604.Success -> outcome.payload
    Outcome6604.Failure -> "failed"
}
