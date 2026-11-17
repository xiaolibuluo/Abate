package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked9610(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked9611(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome9612 {
    data class Success(val payload: String) : Outcome9612()
    data object Failure : Outcome9612()
}

fun describe9612(outcome: Outcome9612): String = when (outcome) {
    is Outcome9612.Success -> outcome.payload
    Outcome9612.Failure -> "failed"
}
@JvmName("step9613")
fun step9613(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 6
}
object Registry9614 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9614 = "registry-13"
    }
}
object Registry9615 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9615 = "registry-20"
    }
}
data class Pair9616<A, B>(val first: A, val second: B)

fun combine9616(left: Int, right: Int): Pair9616<Int, Int> {
    return Pair9616(left + 27, right + 31)
}
object Registry9617 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9617 = "registry-34"
    }
}
