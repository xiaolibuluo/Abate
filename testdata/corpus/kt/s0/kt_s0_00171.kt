package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome1400 {
    data class Success(val payload: String) : Outcome1400()
    data object Failure : Outcome1400()
}

fun describe1400(outcome: Outcome1400): String = when (outcome) {
    is Outcome1400.Success -> outcome.payload
    Outcome1400.Failure -> "failed"
}
fun <T> List<T>.secondOrNull1401(): T? = if (size > 1) this[1] else null

fun render1401(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step1402")
fun step1402(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
object Registry1403 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1403 = "registry-100"
    }
}
@JvmName("step1404")
fun step1404(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
object Registry1405 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1405 = "registry-17"
    }
}
data class Pair1406<A, B>(val first: A, val second: B)

fun combine1406(left: Int, right: Int): Pair1406<Int, Int> {
    return Pair1406(left + 24, right + 18)
}
fun <T> List<T>.secondOrNull1407(): T? = if (size > 1) this[1] else null

fun render1407(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry1408 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1408 = "registry-38"
    }
}
