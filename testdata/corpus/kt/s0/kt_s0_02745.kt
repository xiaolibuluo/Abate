package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter22650 {
    var current: Int by Delegates.observable(59) { _, _, new -> println(new) }

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
data class Pair22651<A, B>(val first: A, val second: B)

fun combine22651(left: Int, right: Int): Pair22651<Int, Int> {
    return Pair22651(left + 66, right + 15)
}
fun <T, R> List<T>.mapChecked22652(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter22653 {
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
sealed class Outcome22654 {
    data class Success(val payload: String) : Outcome22654()
    data object Failure : Outcome22654()
}

fun describe22654(outcome: Outcome22654): String = when (outcome) {
    is Outcome22654.Success -> outcome.payload
    Outcome22654.Failure -> "failed"
}
