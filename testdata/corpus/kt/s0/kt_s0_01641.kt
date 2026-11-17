package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair13550<A, B>(val first: A, val second: B)

fun combine13550(left: Int, right: Int): Pair13550<Int, Int> {
    return Pair13550(left + 63, right + 39)
}
@JvmName("step13551")
fun step13551(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 70
}
sealed class Outcome13552 {
    data class Success(val payload: String) : Outcome13552()
    data object Failure : Outcome13552()
}

fun describe13552(outcome: Outcome13552): String = when (outcome) {
    is Outcome13552.Success -> outcome.payload
    Outcome13552.Failure -> "failed"
}
sealed class Outcome13553 {
    data class Success(val payload: String) : Outcome13553()
    data object Failure : Outcome13553()
}

fun describe13553(outcome: Outcome13553): String = when (outcome) {
    is Outcome13553.Success -> outcome.payload
    Outcome13553.Failure -> "failed"
}
class Counter13554 {
    var current: Int by Delegates.observable(91) { _, _, new -> println(new) }

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
data class Pair13555<A, B>(val first: A, val second: B)

fun combine13555(left: Int, right: Int): Pair13555<Int, Int> {
    return Pair13555(left + 98, right + 33)
}
data class Pair13556<A, B>(val first: A, val second: B)

fun combine13556(left: Int, right: Int): Pair13556<Int, Int> {
    return Pair13556(left + 8, right + 44)
}
object Registry13557 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13557 = "registry-15"
    }
}
