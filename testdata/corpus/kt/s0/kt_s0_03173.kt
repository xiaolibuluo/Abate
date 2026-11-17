package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair26060<A, B>(val first: A, val second: B)

fun combine26060(left: Int, right: Int): Pair26060<Int, Int> {
    return Pair26060(left + 30, right + 12)
}
fun <T, R> List<T>.mapChecked26061(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome26062 {
    data class Success(val payload: String) : Outcome26062()
    data object Failure : Outcome26062()
}

fun describe26062(outcome: Outcome26062): String = when (outcome) {
    is Outcome26062.Success -> outcome.payload
    Outcome26062.Failure -> "failed"
}
object Registry26063 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26063 = "registry-51"
    }
}
class Counter26064 {
    var current: Int by Delegates.observable(58) { _, _, new -> println(new) }

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
