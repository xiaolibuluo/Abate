package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked25560(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair25561<A, B>(val first: A, val second: B)

fun combine25561(left: Int, right: Int): Pair25561<Int, Int> {
    return Pair25561(left + 66, right + 39)
}
data class Pair25562<A, B>(val first: A, val second: B)

fun combine25562(left: Int, right: Int): Pair25562<Int, Int> {
    return Pair25562(left + 73, right + 50)
}
class Counter25563 {
    var current: Int by Delegates.observable(80) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked25564(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome25565 {
    data class Success(val payload: String) : Outcome25565()
    data object Failure : Outcome25565()
}

fun describe25565(outcome: Outcome25565): String = when (outcome) {
    is Outcome25565.Success -> outcome.payload
    Outcome25565.Failure -> "failed"
}
