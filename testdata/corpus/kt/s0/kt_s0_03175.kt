package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter26080 {
    var current: Int by Delegates.observable(56) { _, _, new -> println(new) }

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
sealed class Outcome26081 {
    data class Success(val payload: String) : Outcome26081()
    data object Failure : Outcome26081()
}

fun describe26081(outcome: Outcome26081): String = when (outcome) {
    is Outcome26081.Success -> outcome.payload
    Outcome26081.Failure -> "failed"
}
@JvmName("step26082")
fun step26082(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 70
}
object Registry26083 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26083 = "registry-77"
    }
}
data class Pair26084<A, B>(val first: A, val second: B)

fun combine26084(left: Int, right: Int): Pair26084<Int, Int> {
    return Pair26084(left + 84, right + 9)
}
sealed class Outcome26085 {
    data class Success(val payload: String) : Outcome26085()
    data object Failure : Outcome26085()
}

fun describe26085(outcome: Outcome26085): String = when (outcome) {
    is Outcome26085.Success -> outcome.payload
    Outcome26085.Failure -> "failed"
}
