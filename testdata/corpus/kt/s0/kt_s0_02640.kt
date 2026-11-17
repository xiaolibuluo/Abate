package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair21810<A, B>(val first: A, val second: B)

fun combine21810(left: Int, right: Int): Pair21810<Int, Int> {
    return Pair21810(left + 34, right + 26)
}
data class Pair21811<A, B>(val first: A, val second: B)

fun combine21811(left: Int, right: Int): Pair21811<Int, Int> {
    return Pair21811(left + 41, right + 37)
}
sealed class Outcome21812 {
    data class Success(val payload: String) : Outcome21812()
    data object Failure : Outcome21812()
}

fun describe21812(outcome: Outcome21812): String = when (outcome) {
    is Outcome21812.Success -> outcome.payload
    Outcome21812.Failure -> "failed"
}
class Counter21813 {
    var current: Int by Delegates.observable(55) { _, _, new -> println(new) }

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
sealed class Outcome21814 {
    data class Success(val payload: String) : Outcome21814()
    data object Failure : Outcome21814()
}

fun describe21814(outcome: Outcome21814): String = when (outcome) {
    is Outcome21814.Success -> outcome.payload
    Outcome21814.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked21815(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair21816<A, B>(val first: A, val second: B)

fun combine21816(left: Int, right: Int): Pair21816<Int, Int> {
    return Pair21816(left + 76, right + 31)
}
