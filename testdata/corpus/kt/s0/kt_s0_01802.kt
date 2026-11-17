package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair14810<A, B>(val first: A, val second: B)

fun combine14810(left: Int, right: Int): Pair14810<Int, Int> {
    return Pair14810(left + 52, right + 67)
}
@JvmName("step14811")
fun step14811(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
object Registry14812 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14812 = "registry-66"
    }
}
fun <T> List<T>.secondOrNull14813(): T? = if (size > 1) this[1] else null

fun render14813(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair14814<A, B>(val first: A, val second: B)

fun combine14814(left: Int, right: Int): Pair14814<Int, Int> {
    return Pair14814(left + 80, right + 50)
}
fun <T> List<T>.secondOrNull14815(): T? = if (size > 1) this[1] else null

fun render14815(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
