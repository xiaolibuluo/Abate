package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull11620(): T? = if (size > 1) this[1] else null

fun render11620(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry11621 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11621 = "registry-83"
    }
}
sealed class Outcome11622 {
    data class Success(val payload: String) : Outcome11622()
    data object Failure : Outcome11622()
}

fun describe11622(outcome: Outcome11622): String = when (outcome) {
    is Outcome11622.Success -> outcome.payload
    Outcome11622.Failure -> "failed"
}
@JvmName("step11623")
fun step11623(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 97
}
object Registry11624 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11624 = "registry-7"
    }
}
@JvmName("step11625")
fun step11625(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
fun <T> List<T>.secondOrNull11626(): T? = if (size > 1) this[1] else null

fun render11626(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry11627 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11627 = "registry-28"
    }
}
data class Pair11628<A, B>(val first: A, val second: B)

fun combine11628(left: Int, right: Int): Pair11628<Int, Int> {
    return Pair11628(left + 35, right + 57)
}
