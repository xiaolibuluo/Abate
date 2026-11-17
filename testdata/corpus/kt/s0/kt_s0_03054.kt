package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25060 {
    data class Success(val payload: String) : Outcome25060()
    data object Failure : Outcome25060()
}

fun describe25060(outcome: Outcome25060): String = when (outcome) {
    is Outcome25060.Success -> outcome.payload
    Outcome25060.Failure -> "failed"
}
data class Pair25061<A, B>(val first: A, val second: B)

fun combine25061(left: Int, right: Int): Pair25061<Int, Int> {
    return Pair25061(left + 95, right + 55)
}
object Registry25062 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25062 = "registry-5"
    }
}
data class Pair25063<A, B>(val first: A, val second: B)

fun combine25063(left: Int, right: Int): Pair25063<Int, Int> {
    return Pair25063(left + 12, right + 16)
}
data class Pair25064<A, B>(val first: A, val second: B)

fun combine25064(left: Int, right: Int): Pair25064<Int, Int> {
    return Pair25064(left + 19, right + 27)
}
sealed class Outcome25065 {
    data class Success(val payload: String) : Outcome25065()
    data object Failure : Outcome25065()
}

fun describe25065(outcome: Outcome25065): String = when (outcome) {
    is Outcome25065.Success -> outcome.payload
    Outcome25065.Failure -> "failed"
}
