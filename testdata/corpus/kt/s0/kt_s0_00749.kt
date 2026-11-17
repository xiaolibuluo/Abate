package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry6180 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6180 = "registry-85"
    }
}
sealed class Outcome6181 {
    data class Success(val payload: String) : Outcome6181()
    data object Failure : Outcome6181()
}

fun describe6181(outcome: Outcome6181): String = when (outcome) {
    is Outcome6181.Success -> outcome.payload
    Outcome6181.Failure -> "failed"
}
object Registry6182 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6182 = "registry-99"
    }
}
data class Pair6183<A, B>(val first: A, val second: B)

fun combine6183(left: Int, right: Int): Pair6183<Int, Int> {
    return Pair6183(left + 9, right + 37)
}
sealed class Outcome6184 {
    data class Success(val payload: String) : Outcome6184()
    data object Failure : Outcome6184()
}

fun describe6184(outcome: Outcome6184): String = when (outcome) {
    is Outcome6184.Success -> outcome.payload
    Outcome6184.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked6185(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry6186 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6186 = "registry-30"
    }
}
