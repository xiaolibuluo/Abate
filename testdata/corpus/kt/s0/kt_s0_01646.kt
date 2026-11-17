package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry13590 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13590 = "registry-18"
    }
}
object Registry13591 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13591 = "registry-25"
    }
}
object Registry13592 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13592 = "registry-32"
    }
}
sealed class Outcome13593 {
    data class Success(val payload: String) : Outcome13593()
    data object Failure : Outcome13593()
}

fun describe13593(outcome: Outcome13593): String = when (outcome) {
    is Outcome13593.Success -> outcome.payload
    Outcome13593.Failure -> "failed"
}
object Registry13594 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13594 = "registry-46"
    }
}
data class Pair13595<A, B>(val first: A, val second: B)

fun combine13595(left: Int, right: Int): Pair13595<Int, Int> {
    return Pair13595(left + 53, right + 61)
}
sealed class Outcome13596 {
    data class Success(val payload: String) : Outcome13596()
    data object Failure : Outcome13596()
}

fun describe13596(outcome: Outcome13596): String = when (outcome) {
    is Outcome13596.Success -> outcome.payload
    Outcome13596.Failure -> "failed"
}
class Counter13597 {
    var current: Int by Delegates.observable(67) { _, _, new -> println(new) }

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
