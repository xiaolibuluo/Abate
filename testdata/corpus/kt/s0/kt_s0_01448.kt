package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11950<A, B>(val first: A, val second: B)

fun combine11950(left: Int, right: Int): Pair11950<Int, Int> {
    return Pair11950(left + 20, right + 17)
}
sealed class Outcome11951 {
    data class Success(val payload: String) : Outcome11951()
    data object Failure : Outcome11951()
}

fun describe11951(outcome: Outcome11951): String = when (outcome) {
    is Outcome11951.Success -> outcome.payload
    Outcome11951.Failure -> "failed"
}
fun <T> List<T>.secondOrNull11952(): T? = if (size > 1) this[1] else null

fun render11952(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair11953<A, B>(val first: A, val second: B)

fun combine11953(left: Int, right: Int): Pair11953<Int, Int> {
    return Pair11953(left + 41, right + 50)
}
object Registry11954 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11954 = "registry-48"
    }
}
data class Pair11955<A, B>(val first: A, val second: B)

fun combine11955(left: Int, right: Int): Pair11955<Int, Int> {
    return Pair11955(left + 55, right + 11)
}
