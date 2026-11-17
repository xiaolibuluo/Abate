package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25130 {
    data class Success(val payload: String) : Outcome25130()
    data object Failure : Outcome25130()
}

fun describe25130(outcome: Outcome25130): String = when (outcome) {
    is Outcome25130.Success -> outcome.payload
    Outcome25130.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked25131(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair25132<A, B>(val first: A, val second: B)

fun combine25132(left: Int, right: Int): Pair25132<Int, Int> {
    return Pair25132(left + 96, right + 54)
}
data class Pair25133<A, B>(val first: A, val second: B)

fun combine25133(left: Int, right: Int): Pair25133<Int, Int> {
    return Pair25133(left + 6, right + 65)
}
class Counter25134 {
    var current: Int by Delegates.observable(13) { _, _, new -> println(new) }

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
object Registry25135 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25135 = "registry-20"
    }
}
