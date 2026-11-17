/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

// ktRecipe covers packages, sealed/data classes, generics, extension
// functions, when expressions, lambdas, property delegation, suspend functions
// and null-safety operators.
func ktRecipe() Recipe {
	return Recipe{
		Language: "kt",
		Ext:      "kt",
		Tags:     []string{"class", "generic", "extension", "when", "lambda", "null-safety", "exception"},
		Header:   "package synth\n\nimport kotlin.properties.Delegates\n\n",
		AnchorClean: `fun anchorProbe(): Boolean {
    return 1 + 2 == 3
}

`,
		AnchorFull: `fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

`,
		CleanBlock: "fun step@@I@@x@@S@@(): Int = @@N@@ + @@S@@\n",
		Blocks: []BlockTemplate{
			{Text: `sealed class Outcome@@I@@@@S@@ {
    data class Success(val payload: String) : Outcome@@I@@@@S@@()
    data object Failure : Outcome@@I@@@@S@@()
}

fun describe@@I@@@@S@@(outcome: Outcome@@I@@@@S@@): String = when (outcome) {
    is Outcome@@I@@@@S@@.Success -> outcome.payload
    Outcome@@I@@@@S@@.Failure -> "failed"
}
`},
			{Text: `fun <T> List<T>.secondOrNull@@I@@@@S@@(): T? = if (size > 1) this[1] else null

fun render@@I@@@@S@@(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
`},
			{Text: `class Counter@@I@@@@S@@ {
    var current: Int by Delegates.observable(@@N@@) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
`},
			{Text: `object Registry@@I@@@@S@@ {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG@@I@@@@S@@ = "registry-@@N@@"
    }
}
`},
			{Text: `@JvmName("step@@I@@@@S@@")
fun step@@I@@@@S@@(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + @@N@@
}
`},
			{Text: `fun <T, R> List<T>.mapChecked@@I@@@@S@@(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
`},
			{Text: `data class Pair@@I@@@@S@@<A, B>(val first: A, val second: B)

fun combine@@I@@@@S@@(left: Int, right: Int): Pair@@I@@@@S@@<Int, Int> {
    return Pair@@I@@@@S@@(left + @@N@@, right + @@M@@)
}
`},
		},
	}
}
