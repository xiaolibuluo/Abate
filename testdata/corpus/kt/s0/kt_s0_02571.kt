package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked21180(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair21181<A, B>(val first: A, val second: B)

fun combine21181(left: Int, right: Int): Pair21181<Int, Int> {
    return Pair21181(left + 95, right + 23)
}
sealed class Outcome21182 {
    data class Success(val payload: String) : Outcome21182()
    data object Failure : Outcome21182()
}

fun describe21182(outcome: Outcome21182): String = when (outcome) {
    is Outcome21182.Success -> outcome.payload
    Outcome21182.Failure -> "failed"
}
object Registry21183 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21183 = "registry-12"
    }
}
data class Pair21184<A, B>(val first: A, val second: B)

fun combine21184(left: Int, right: Int): Pair21184<Int, Int> {
    return Pair21184(left + 19, right + 56)
}
sealed class Outcome21185 {
    data class Success(val payload: String) : Outcome21185()
    data object Failure : Outcome21185()
}

fun describe21185(outcome: Outcome21185): String = when (outcome) {
    is Outcome21185.Success -> outcome.payload
    Outcome21185.Failure -> "failed"
}
