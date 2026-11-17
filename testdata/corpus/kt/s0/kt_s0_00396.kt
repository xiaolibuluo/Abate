package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3250<A, B>(val first: A, val second: B)

fun combine3250(left: Int, right: Int): Pair3250<Int, Int> {
    return Pair3250(left + 59, right + 27)
}
fun <T> List<T>.secondOrNull3251(): T? = if (size > 1) this[1] else null

fun render3251(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry3252 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3252 = "registry-73"
    }
}
@JvmName("step3253")
fun step3253(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
sealed class Outcome3254 {
    data class Success(val payload: String) : Outcome3254()
    data object Failure : Outcome3254()
}

fun describe3254(outcome: Outcome3254): String = when (outcome) {
    is Outcome3254.Success -> outcome.payload
    Outcome3254.Failure -> "failed"
}
object Registry3255 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3255 = "registry-94"
    }
}
data class Pair3256<A, B>(val first: A, val second: B)

fun combine3256(left: Int, right: Int): Pair3256<Int, Int> {
    return Pair3256(left + 101, right + 32)
}
