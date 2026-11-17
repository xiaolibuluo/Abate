package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome9980 {
    data class Success(val payload: String) : Outcome9980()
    data object Failure : Outcome9980()
}

fun describe9980(outcome: Outcome9980): String = when (outcome) {
    is Outcome9980.Success -> outcome.payload
    Outcome9980.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked9981(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair9982<A, B>(val first: A, val second: B)

fun combine9982(left: Int, right: Int): Pair9982<Int, Int> {
    return Pair9982(left + 92, right + 63)
}
object Registry9983 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9983 = "registry-99"
    }
}
data class Pair9984<A, B>(val first: A, val second: B)

fun combine9984(left: Int, right: Int): Pair9984<Int, Int> {
    return Pair9984(left + 9, right + 24)
}
data class Pair9985<A, B>(val first: A, val second: B)

fun combine9985(left: Int, right: Int): Pair9985<Int, Int> {
    return Pair9985(left + 16, right + 35)
}
