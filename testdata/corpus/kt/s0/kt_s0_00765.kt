package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome6320 {
    data class Success(val payload: String) : Outcome6320()
    data object Failure : Outcome6320()
}

fun describe6320(outcome: Outcome6320): String = when (outcome) {
    is Outcome6320.Success -> outcome.payload
    Outcome6320.Failure -> "failed"
}
fun <T> List<T>.secondOrNull6321(): T? = if (size > 1) this[1] else null

fun render6321(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry6322 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6322 = "registry-87"
    }
}
fun <T> List<T>.secondOrNull6323(): T? = if (size > 1) this[1] else null

fun render6323(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step6324")
fun step6324(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
@JvmName("step6325")
fun step6325(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
object Registry6326 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6326 = "registry-18"
    }
}
object Registry6327 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6327 = "registry-25"
    }
}
data class Pair6328<A, B>(val first: A, val second: B)

fun combine6328(left: Int, right: Int): Pair6328<Int, Int> {
    return Pair6328(left + 32, right + 68)
}
sealed class Outcome6329 {
    data class Success(val payload: String) : Outcome6329()
    data object Failure : Outcome6329()
}

fun describe6329(outcome: Outcome6329): String = when (outcome) {
    is Outcome6329.Success -> outcome.payload
    Outcome6329.Failure -> "failed"
}
