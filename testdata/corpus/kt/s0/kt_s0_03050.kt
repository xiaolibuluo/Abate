package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry25030 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25030 = "registry-49"
    }
}
fun <T> List<T>.secondOrNull25031(): T? = if (size > 1) this[1] else null

fun render25031(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair25032<A, B>(val first: A, val second: B)

fun combine25032(left: Int, right: Int): Pair25032<Int, Int> {
    return Pair25032(left + 63, right + 45)
}
data class Pair25033<A, B>(val first: A, val second: B)

fun combine25033(left: Int, right: Int): Pair25033<Int, Int> {
    return Pair25033(left + 70, right + 56)
}
sealed class Outcome25034 {
    data class Success(val payload: String) : Outcome25034()
    data object Failure : Outcome25034()
}

fun describe25034(outcome: Outcome25034): String = when (outcome) {
    is Outcome25034.Success -> outcome.payload
    Outcome25034.Failure -> "failed"
}
