package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7210 {
    data class Success(val payload: String) : Outcome7210()
    data object Failure : Outcome7210()
}

fun describe7210(outcome: Outcome7210): String = when (outcome) {
    is Outcome7210.Success -> outcome.payload
    Outcome7210.Failure -> "failed"
}
sealed class Outcome7211 {
    data class Success(val payload: String) : Outcome7211()
    data object Failure : Outcome7211()
}

fun describe7211(outcome: Outcome7211): String = when (outcome) {
    is Outcome7211.Success -> outcome.payload
    Outcome7211.Failure -> "failed"
}
class Counter7212 {
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
sealed class Outcome7213 {
    data class Success(val payload: String) : Outcome7213()
    data object Failure : Outcome7213()
}

fun describe7213(outcome: Outcome7213): String = when (outcome) {
    is Outcome7213.Success -> outcome.payload
    Outcome7213.Failure -> "failed"
}
data class Pair7214<A, B>(val first: A, val second: B)

fun combine7214(left: Int, right: Int): Pair7214<Int, Int> {
    return Pair7214(left + 94, right + 37)
}
fun <T, R> List<T>.mapChecked7215(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry7216 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7216 = "registry-11"
    }
}
